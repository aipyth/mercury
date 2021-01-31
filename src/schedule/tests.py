from django.test import TestCase
from datetime import time
import json
from .models import TimeSheme
from .serializers import TimeShemeSerializer
from rest_framework import serializers


class TimeSchemeTestCase(TestCase):

    pass


class TimeShemeSerializerTest(TestCase):
    def test_to_representation(self):
        ts = TimeSheme.objects.create(name='test', items=json.dumps([
            {'start': time(hour=13).isoformat(),
             'stop': time(hour=14).isoformat()},
            {'start': time(hour=15).isoformat(),
             'stop': time(hour=16).isoformat()}
        ]))
        tss = TimeShemeSerializer(ts)
        self.assertTrue(isinstance(tss.data['items'],
                        list))
