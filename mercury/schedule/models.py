from django.db import models
from django.contrib.auth.models import User


class Room(models.Model):
    name = models.CharField(max_length=200,
                            blank=True, null=True)
    owner = models.ForeignKey(User, on_delete=models.CASCADE)
    period = models.IntegerField()
    start_date = models.DateField()
    end_date = models.DateField()
    public = models.BooleanField(default=False)


class Subject(models.Model):
    room = models.ForeignKey(Room, on_delete=models.CASCADE,
                             related_name='subjects')
    name = models.CharField(max_length=200)
    lector = models.CharField(max_length=200, blank=True, null=True)
    day = models.IntegerField()
    starts = models.TimeField(null=True)
    duration = models.TimeField(null=True)
