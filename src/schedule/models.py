from django.db import models
# from django.contrib.auth.models import User
from django.conf import settings


class TimeSheme(models.Model):
    name = models.CharField(max_length=200)
    items = models.TextField()

    created = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.name}"

    def get_pythonized_items(self):
        pass


class Room(models.Model):
    name = models.CharField(max_length=200,
                            blank=True, null=True)
    owner = models.ForeignKey(settings.AUTH_USER_MODEL,
                              on_delete=models.CASCADE)
    period = models.IntegerField()
    start_date = models.DateField()
    end_date = models.DateField()
    public = models.BooleanField(default=False)

    created = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.name} by {self.owner}"


class Subject(models.Model):
    room = models.ForeignKey(Room, on_delete=models.CASCADE,
                             related_name='subjects')
    name = models.CharField(max_length=200)
    days_and_orders = models.TextField(default="[]")
    lector = models.CharField(max_length=200, blank=True, null=True)
    extra = models.TextField(blank=True, null=True)

    def __str__(self):
        return f"{self.name}"
