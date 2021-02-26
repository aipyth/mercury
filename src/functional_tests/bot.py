import pyrogram
import unittest
import os

BOT_USERNAME = os.environ.get('BOT_USERNAME', 'AthenaMercuryBot')
API_ID = os.environ.get('APP_ID')
API_HASH = os.environ.get('APP_HASH')


class InlineGetScheduleTests(unittest.TestCase):
    @classmethod
    def setUpClass(self):
        self.app = pyrogram.Client('testcaseclient', API_ID, API_HASH)
        self.app.start()

    @classmethod
    def tearDownClass(self):
        self.app.stop()

    def test_inline_result_not_empty(self):
        response = self.app.get_inline_bot_results(BOT_USERNAME, 'test-room')
        self.assertNotEqual([], response.results)
        # response.results


if __name__ == '__main__':
    unittest.main(warnings='ignore')
