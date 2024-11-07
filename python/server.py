from flask import Flask, request, jsonify
from langchain_community.document_loaders import PyPDFLoader
from langchain_community.chat_models import ChatOpenAI
from langchain.chains.question_answering import load_qa_chain
import json
import psycopg2

app = Flask(__name__)

@app.route('/stock/reports', methods=['POST'])
def ask_question():
    # Get the request data
    data = request.get_json()
    pdf_path = data.get('file_path')
    question = """
Can you analyze the provided stock research report and extract relevant data to populate a StockReport object? Please ensure to follow these instructions:

1. Extract only verified and factual information directly from the report—do not infer or assume any details.

2. For the report, please fill out the following StockReport struct:

type StockReport struct {
    Company           string   `json:"company"`              // Name of the company
    Ticker            string   `json:"ticker"`              // Ticker name of company (fetch from NSE India if not present)
    Sector            string   `json:"sector"`               // Industry sector
    Recommendation    string   `json:"recommendation"`       // Analyst recommendation
    TargetPrice       float64  `json:"target_price"`         // Target price in INR
    RevenueProjections map[string]float64 `json:"revenue_projections"` // Revenue projections
    CAGR              float64  `json:"cagr"`                 // CAGR
    EBITDA            map[string]float64 `json:"ebitda"`         // EBITDA values
    Date              string   `json:"date"`                 // Date of the report
    NewsSummary       string   `json:"news_summary"`         // Summary of key news
}

3. Guidance on extraction:
   - Company name and sector are typically found in the title or introductory section of the report.
   - Recommendation and target price are located on the first or second page under a headline.
   - Revenue projections are found in the financial tables alongside future fiscal years.
   - CAGR and EBITDA values are found in the financial projections section.
   - News summary is extracted from narrative sections detailing company updates or significant events (around 150 words).

4. Cross-check each extracted field for accuracy. If data is missing, mark it as null in the JSON output.

5. The final output must be formatted as valid JSON, containing only one StockReport object in a single line, with no extra whitespace or newline characters.

{
  "data": {
    "company": "string",
    "ticker": "string", //Fetch ticker associated with the company from NSE India
    "sector": "string",
    "recommendation": "string",
    "target_price": "float64",
    "revenue_projections": {"FY24": value, "FY25": value},
    "cagr": "float64",
    "ebitda": {"FY24": value, "FY25": value},
    "date": "01/01/2023",
    "news_summary": "string"
  },
  "err": null if successful, or a relevant error message if extraction fails.
}

6. The 'err' object in the response is required to be null if data is successfully fetched or contain a relevant error message if extraction fails.

7. IMPORTANT: Only process the relevant stock research report provided.

8. IMPORTANT: Return the output as valid JSON in a single line without any additional formatting, including no newline characters.
9. If the ticker is not present in the report, please fetch the ticker associated with the company name from NSE India.

"""


    if not pdf_path or not question:
        return jsonify({'error': 'file_path and question are required.'}), 400

    try:
        # Load the PDF
        loader = PyPDFLoader(pdf_path)
        documents = loader.load()

        # Extract only the first two pages
        if len(documents) > 2:
            documents = documents[:2]  # Keep only the first two pages

        # Initialize the OpenAI API
        api_key = "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"  # Replace with your OpenAI API key
        llm = ChatOpenAI(model_name="chatgpt-4o-latest",openai_api_key=api_key)
        chain = load_qa_chain(llm, verbose=True)

        # Get the response
        response = chain.run(input_documents=documents, question=question)
        process_and_insert(response)
        parsed_response = json.loads(response)
        return parsed_response, 200

    except Exception as e:
        return jsonify({'error': str(e)}), 500

def process_and_insert(response):
    try:
        # Parse the JSON response
        parsed_response = json.loads(response)
        
        # Check for error in the response
        if parsed_response.get("err") is not None:
            print(f"Error in response: {parsed_response['err']}")
            return
        
        # Extract the 'data' field
        data = parsed_response.get("data")
        
        # Insert the extracted data into the stock_reports table
        insert_stock_report(data)

    except json.JSONDecodeError as e:
        print(f"Error parsing JSON: {e}")


db_params = {
    'dbname': 'chat',
    'user': 'nexbit',
    'password': 'password',
    'host': 'localhost',  # or your database host
    'port': '5432'        # default PostgreSQL port
}
	
# Function to insert data into PostgreSQL database
def insert_stock_report(data):
    try:
        # Connect to PostgreSQL database
        conn = psycopg2.connect(**db_params)
        cursor = conn.cursor()

        # Define the insert SQL statement
        insert_query = """
        INSERT INTO stock_reports (company, sector, recommendation, target_price, revenue_projections, cagr, ebitda, news_summary, date,ticker)
        VALUES (%s, %s, %s, %s, %s::jsonb, %s, %s::jsonb, %s, %s,%s)
        RETURNING id;
        """
        
        # Execute the insert statement
        cursor.execute(insert_query, (
            data['company'],
            data['sector'],
            data['recommendation'],
            data['target_price'],
            json.dumps(data['revenue_projections']),  # Convert dict to JSON string
            data['cagr'],
            json.dumps(data['ebitda']),               # Convert dict to JSON string
            data['news_summary'],
            data['date'],
            data['ticker']
        ))

        # Fetch the returned id
        new_id = cursor.fetchone()[0]
        print(f"Inserted stock report with ID: {new_id}")

        # Commit the transaction
        conn.commit()

    except Exception as e:
        print(f"Error inserting data: {e}")
    finally:
        # Close cursor and connection
        cursor.close()
        conn.close()

 
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3004)


