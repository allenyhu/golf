import gspread
from gspread_formatting import *

client = gspread.oauth()
sh = client.open("Golf Tracker").worksheet('Penmar 12/01/24')
format_range = "E:E"


current_rules = get_conditional_format_rules(sh)
current_rules.extend(new_rules)

current_rules.save()


def create_putting_rules():
    return [
        ConditionalFormatRule(
            ranges=[GridRange.from_a1_range(format_range, sh)],
            booleanRule=BooleanRule(
                condition=BooleanCondition('NUMBER_GREATER', ['2']),
                format=CellFormat(backgroundColor=Color(1, 0, 0))  # Red
            )
        ),
        ConditionalFormatRule(
            ranges=[GridRange.from_a1_range(format_range, sh)],
            booleanRule=BooleanRule(
                condition=BooleanCondition('NUMBER_LESS', ['2']),
                format=CellFormat(backgroundColor=Color(0, 1, 0))  # Green
            )
        )
    ]
