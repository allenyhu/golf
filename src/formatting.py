import gspread
from gspread_formatting import *
from google.oauth2.credentials import Credentials
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError


def create_threshold_formatting(sheet, cell_range, threshold):
    return [
        ConditionalFormatRule(
            ranges=[GridRange.from_a1_range(cell_range, sheet)],
            booleanRule=BooleanRule(
                condition=BooleanCondition('NUMBER_GREATER', [threshold]),
                format=CellFormat(backgroundColor=Color(1, 0, 0))  # Red
            )
        ),
        ConditionalFormatRule(
            ranges=[GridRange.from_a1_range(cell_range, sheet)],
            booleanRule=BooleanRule(
                condition=BooleanCondition('NUMBER_LESS', [threshold]),
                format=CellFormat(backgroundColor=Color(0, 1, 0))  # Green
            )
        )
    ]

def get_sheet_rules(sheet):
    """Creates and returns formatting rules for the golf score sheet"""
    rules = [
        # Make first row bold
        ConditionalFormatRule(
            ranges=[GridRange.from_a1_range("A1:I1", sheet)],
            booleanRule=BooleanRule(
                condition=BooleanCondition('CUSTOM_FORMULA', ['=ROW()=1']),
                format=CellFormat(textFormat=TextFormat(bold=True))
            )
        )
    ]
    
    # Add score difference formatting (column D)
    score_rules = create_threshold_formatting(sheet, "D:D", "2")
    rules.extend(score_rules)
    
    # Add putting formatting (column E) 
    putting_rules = create_threshold_formatting(sheet, "E:E", "2")
    rules.extend(putting_rules)

    chipping_rules = create_threshold_formatting(sheet, "G:G", "1")
    rules.extend(chipping_rules)
    
    # Add strokes-to-green formatting (column H)
    # Use par value from column B as threshold
    stg_rules = create_threshold_formatting(sheet, "H:H", "=B:B")  # References par value
    rules.extend(stg_rules)
    
    return rules

def format_sheet(sheet_id):
    """Apply formatting using Google Sheets API v4"""
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
                }
            },
            # Score difference formatting (column D)
            create_threshold_rule(sheet_id, 3, 3, "2"),  # Column D (0-based index 3)
            # Putting formatting (column E)
            create_threshold_rule(sheet_id, 4, 4, "2"),  # Column E
            # Chipping formatting (column G)
            create_threshold_rule(sheet_id, 6, 6, "1"),  # Column G
            # STG formatting (column H)
            create_threshold_rule(sheet_id, 7, 7, "=B:B")  # Column H
        ]
    }
    return requests

def create_threshold_rule(sheet_id, start_col, end_col, threshold):
    """Creates threshold-based conditional formatting rules"""
    return {
        "addConditionalFormatRule": {
            "rule": {
                "ranges": [{
                    "sheetId": sheet_id,
                    "startColumnIndex": start_col,
                    "endColumnIndex": end_col + 1
                }],
                "booleanRule": {
                    "condition": {
                        "type": "NUMBER_GREATER",
                        "values": [{"userEnteredValue": threshold}]
                    },
                    "format": {
                        "backgroundColor": {"red": 1, "green": 0, "blue": 0}
                    }
                }
            }
        }
    }

