
def format_sheet(sheet_id):
    return {
        'requests': [
            {
                'addConditionalFormatRule': create_putting_rule(sheet_id)
                
            }
        ]
    }


def create_putting_rule(sheet_id):
    return {
        'rule': {
            'ranges': [
                {
                    'sheetId': sheet_id,
                    'startColumnIndex': 4,
                    'endColumnIndex': 5,
                }
            ],
            'booleanRule': {
                'condition': {
                    'type': 'NUMBER_GREATER',
                    'values': [
                        {
                            'userEnteredValue': '2'
                        }
                    ]
                },
                'format': {
                    'backgroundColor': {
                        'red': 1.0,
                        'green': 0,
                        'blue': 0
                    }
                }
            }
        }}
