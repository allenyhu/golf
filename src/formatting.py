import gspread
from gspread_formatting import *


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

def format_sheet(sheet_name):
    sheet = gspread.oauth().open("Golf Tracker").worksheet(sheet_name)
    
    # Get new rules
    new_rules = get_sheet_rules(sheet)
    
    # Get existing rules, extend with new rules and save
    sheet_rules = get_conditional_format_rules(sheet)
    sheet_rules.extend(new_rules)
    sheet_rules.save()

