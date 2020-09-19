from django.contrib.auth.models import User
from rest_framework import serializers

from schedule.models import Room, Subject


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('username', 'password')
        extra_kwargs = {'password': {'write_only': True}}

    def create(self, validated_data):
        password = validated_data.pop('password')
        user = User(**validated_data)
        user.set_password(password)
        user.save()
        return user


class RoomSerializer(serializers.HyperlinkedModelSerializer):
    subjects = serializers.HyperlinkedRelatedField(
        many=True, read_only=True, view_name='subject-detail'
    )

    class Meta:
        model = Room
        fields = ['url', 'name', 'period', 'start_date', 'end_date', 'public',
                  'subjects']


class SubjectSerializer(serializers.ModelSerializer):
    class Meta:
        model = Subject
        fields = ['name', 'lector', 'room', 'day', 'starts', 'duration']
