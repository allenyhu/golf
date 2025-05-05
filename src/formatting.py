import gspread
from gspread_formatting import *
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError

def format_sheet(sheet_id):
    """Apply formatting using Google Sheets API v4"""
    # Column rules with (column_number, threshold)
    column_rules = [
        (4, "1"),    # Score difference (column D)
        (5, "2"),    # Putting (column E)
        (7, "1"),    # Chipping (column G)
        (8, "=B:B-1")  # Strokes-to-green (column H)
    ]
    
    requests = {
        "requests": [
            {
                "addConditionalFormatRule": {
                    "rule": {
                        "ranges": [{"sheetId": sheet_id, "startRowIndex": 0, "endRowIndex": 1}],
                        "booleanRule": {
                            "condition": {
                                "type": "CUSTOM_FORMULA",
                                "values": [{"userEnteredValue": "=ROW()=1"}]
                            },
                            "format": {
                                "textFormat": {"bold": True}
                            }
                        }
                    }
                },
            },
            {
                "repeatCell": {
                    "range": {
                        "sheetId": sheet_id,
                        "startColumnIndex": 3,  # Column D
                        "endColumnIndex": 4,
                        "startRowIndex": 1,
                        "endRowIndex": 19
                    },
                    "cell": {
                        "userEnteredValue": {
                            "formulaValue": "=C:C-B:B"
                        },
                        "userEnteredFormat": {
                            "numberFormat": {
                                "type": "NUMBER"
                            }
                        }
                    },
                    "fields": "userEnteredValue,userEnteredFormat.numberFormat"
                }
            },
            {
                "repeatCell": {
                    "range": {
                        "sheetId": sheet_id,
                        "startColumnIndex": 7,  # Column H
                        "endColumnIndex": 8,
                        "startRowIndex": 1,
                        "endRowIndex": 19
                    },
                    "cell": {
                        "userEnteredValue": {
                            "formulaValue": "=C:C-E:E-I:I"
                        },
                        "userEnteredFormat": {
                            "numberFormat": {
                                "type": "NUMBER"
                            }
                        }
                    },
                    "fields": "userEnteredValue,userEnteredFormat.numberFormat"
                }
            }
        ]
    }
    
    # Add threshold rules for each column
    for col_num, threshold in column_rules:
        rules = create_threshold_rule(sheet_id, col_num, threshold)
        requests["requests"].extend(rules)
    
    return requests


def create_threshold_rule(sheet_id, col_num, threshold):
    """Creates threshold-based conditional formatting rules for API v4
    
    Args:
        sheet_id: The sheet ID
        col_num: Column number (1-based index)
        threshold: Threshold value or formula for comparison
    """
    col_index = col_num - 1  # Convert to 0-based index
    return [
        {
            "addConditionalFormatRule": {
                "rule": {
                    "ranges": [{
                        "sheetId": sheet_id,
                        "startColumnIndex": col_index,
                        "endColumnIndex": col_index + 1,
                        "startRowIndex": 1
                    }],
                    "booleanRule": {
                        "condition": {
                            "type": "NUMBER_GREATER",
                            "values": [{"userEnteredValue": threshold}]
                        },
                        "format": {
                            "backgroundColor": {"red": 1, "green": 0, "blue": 0}  # Red
                        }
                    }
                }
            }
        },
        {
            "addConditionalFormatRule": {
                "rule": {
                    "ranges": [{
                        "sheetId": sheet_id,
                        "startColumnIndex": col_index,
                        "endColumnIndex": col_index + 1
                    }],
                    "booleanRule": {
                        "condition": {
                            "type": "NUMBER_LESS",
                            "values": [{"userEnteredValue": threshold}]
                        },
                        "format": {
                            "backgroundColor": {"red": 0, "green": 1, "blue": 0}  # Green
                        }
                    }
                }
            }
        }
    ]
