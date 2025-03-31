import formatting
import os.path
import sys

from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError

SCOPES = ["https://www.googleapis.com/auth/spreadsheets"]

GOLF_TRACKER_SHEET_ID = "1R_HHJRloAsyP7BuQFoxtXIzBhFHmTd5M8CiNr-QzccY"
COLUMNS = ['Hole', 'Par', 'Score', 'Diff', 'Putts', 'Tee Shot', 'Chips', 'STG', 'Penalty']


def get_credentials():
    creds = None
    # The file token.json stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists("token.json"):
        creds = Credentials.from_authorized_user_file("token.json", SCOPES)
    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
                "credentials.json", SCOPES
            )
            creds = flow.run_local_server(port=0)
        # Save the credentials for the next run
        with open("token.json", "w") as token:
            token.write(creds.to_json())

    return creds


def main():
    creds = get_credentials()
    sheet_name = sys.argv[1] if len(sys.argv) > 1 else 'Test Sheet'

    try:
        service = build("sheets", "v4", credentials=creds)
        
        # Get the sheet ID from create_new_sheet
        sheet_id = create_new_sheet(service, sheet_name)
        sheet_data = prepare_sheet_data()
        
        # Upload data first
        upload_sheet_data(service, sheet_name, sheet_data)
        
        # Then apply formatting using the numeric sheet_id
        apply_sheet_formatting(service, int(sheet_id))  # Convert to int to ensure it's a number

    except HttpError as err:
        print(err)

def create_new_sheet(service, sheet_name):
    """Creates a new sheet and returns its ID"""
    response = service.spreadsheets().batchUpdate(spreadsheetId=GOLF_TRACKER_SHEET_ID, body={
        "requests": {
            "addSheet": {
                "properties": {
                    "title": sheet_name,
                    "index": 1
                },
            }
        }
    }).execute()
    
    return response['replies'][0]['addSheet']['properties']['sheetId']

def prepare_sheet_data():
    """Prepares all data for the sheet in memory"""
    # Start with the column headers
    data = [COLUMNS]
    
    # Add hole numbers (1-18)
    for i in range(1, 19):
        row = [str(i)] + [''] * (len(COLUMNS) - 1)  # Empty cells for other columns
        data.append(row)

    # Add summary rows
    summary_rows = [
        ["Score", "=SUM(C2:C19)"],
        ["Extra Putts", "=SUM(FILTER(E2:E19 - 2, E2:E19 - 2 >= 0))"],
        ["Extra Chips", "=SUM(FILTER(G2:G19 - 1, G2:G19 - 1 >= 0))"],
        ["Pars", "=COUNTIF(D2:D19,\"<=0\")"],
    ]
    
    # Add empty row as separator
    data.append(["", ""])
    data.append(["", ""])
    
    # Add all summary rows
    for row in summary_rows:
        data.append(row + [""] * (len(COLUMNS) - 2))  # Pad remaining columns with empty strings
        
    return data

def upload_sheet_data(service, sheet_name, data):
    """Uploads all sheet data in a single operation"""
    service.spreadsheets().values().update(
        spreadsheetId=GOLF_TRACKER_SHEET_ID,
        range=f"{sheet_name}!A1",
        valueInputOption="USER_ENTERED",
        body={'values': data}
    ).execute()

def apply_sheet_formatting(service, sheet_id):
    """Applies the predefined formatting to the sheet"""
    format_requests = formatting.format_sheet(sheet_id)
    service.spreadsheets().batchUpdate(
        spreadsheetId=GOLF_TRACKER_SHEET_ID,
        body=format_requests
    ).execute()

if __name__ == "__main__":
    main()
