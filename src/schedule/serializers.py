import json
from rest_framework import serializers

from users.models import CustomUser

from schedule.models import TimeSheme, Room, Subject


class TimeShemeSerializer(serializers.ModelSerializer):
    items = serializers.JSONField()

    class Meta:
        model = TimeSheme
        fields = ('id', 'name', 'items')

    def to_representation(self, instance):
        ret = super().to_representation(instance)
        ret['items'] = json.loads(ret['items'])
        return ret
  
    def to_internal_value(self, data):
        ret = super().to_internal_value(data)
        ret['items'] = json.dumps(ret['items'])
        return ret


class CustomUserSerializer(serializers.ModelSerializer):
    class Meta:
        model = CustomUser
        fields = ('email', 'password')
        extra_kwargs = {'password': {'write_only': True}}

    def create(self, validated_data):
        password = validated_data.pop('password')
        user = CustomUser(**validated_data)
        user.set_password(password)
        user.save()
        return user


# class RoomSerializer(serializers.HyperlinkedModelSerializer):
class RoomSerializer(serializers.ModelSerializer):
    # subjects = serializers.HyperlinkedRelatedField(
    #     many=True, read_only=True, view_name='subject-detail'
    # )

    class Meta:
        model = Room
        exclude = ['owner']
        depth = 1


class SubjectSerializer(serializers.ModelSerializer):
    days_and_orders = serializers.JSONField()
    
    class Meta:
        model = Subject
        fields = '__all__'

    def to_representation(self, instance):
        ret = super().to_representation(instance)
        ret['days_and_orders'] = json.loads(ret['days_and_orders'])
        return ret

    def to_internal_value(self, data):
        ret = super().to_internal_value(data)
        ret['days_and_orders'] = json.dumps(ret['days_and_orders'])
        return ret
