from django.test import TestCase
from datetime import time
import json
from .models import TimeSchema
from .serializers import TimeSchemaSerializer
from rest_framework import serializers


class TimeSchemaTestCase(TestCase):

    pass


class TimeSchemaSerializerTest(TestCase):
    def test_to_representation(self):
        ts = TimeSchema.objects.create(name='test', items=json.dumps([
            {'start': time(hour=13).isoformat(),
             'stop': time(hour=14).isoformat()},
            {'start': time(hour=15).isoformat(),
             'stop': time(hour=16).isoformat()}
        ]))
        tss = TimeSchemaSerializer(ts)
        self.assertTrue(isinstance(tss.data['items'],
                        list))
