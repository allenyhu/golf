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

client = gspread.oauth()
sh = client.open("Golf Tracker").worksheet('Penmar 11/01/24')

sheet_rules = get_conditional_format_rules(sh)

putting_rules = create_putting_rules(sheet, "E:E", "2")
sheet_rules.extend(putting_rules)

sheet_rules.save()