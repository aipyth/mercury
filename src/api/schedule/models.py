from django.db import models
from django.core.exceptions import ObjectDoesNotExist
# from django.contrib.auth.models import User
from django.conf import settings
from contextlib import suppress
from django.core.files.base import ContentFile
from django.core.files.uploadedfile import InMemoryUploadedFile

from contextlib import suppress

import datetime as dt

import imgkit
from PIL import Image
import io
import time
import json


class CustomManager(models.Manager):
    def get_or_none(self, *args, **kwargs):
        with suppress(ObjectDoesNotExist):
            return super().get(*args, **kwargs)
        return None


class TimeSchema(models.Model):
    name = models.CharField(max_length=200)
    items = models.TextField()
    public = models.BooleanField(default=False)

    created = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.name}"

    def get_pythonized_items(self):
        return json.loads(self.items)

    def entries(self):
        items = json.loads(self.items)
        return len(items)


class Room(models.Model):
    name = models.CharField(max_length=200,
                            blank=True, null=True)
    slug = models.SlugField(null=True, unique=True)
    owner = models.ForeignKey(settings.AUTH_USER_MODEL,
                              on_delete=models.CASCADE)
    period = models.IntegerField()
    time_schema = models.ForeignKey(TimeSchema, on_delete=models.CASCADE,
                                   related_name='rooms')
    start_date = models.DateField()
    end_date = models.DateField()
    public = models.BooleanField(default=False)

    schedule_image = models.ImageField(
        upload_to="schedules",
        default='',
        blank=True,
        )
    schedule_image_thumb = models.ImageField(
        upload_to="schedules",
        default='',
        blank=True,
        )

    created = models.DateTimeField(auto_now_add=True)

    objects = CustomManager()

    def __str__(self):
        return f"{self.name} by {self.owner}"


    def generate_schedule_image(self):
        if self.schedule_image != None:
            with suppress(ValueError):
                storage, path = self.schedule_image.storage, self.schedule_image.path
                storage.delete(path)
        url = "http://web:8000/room/" + str(self.id) + "/"
        image_data = imgkit.from_url(url, False, options={
            "width": 1000,
            "height": 1000,
        })
        image_data = io.BytesIO(image_data)
        im = ContentFile(image_data.getvalue())
        image_name = str(self.id) + ".jpg"
        self.schedule_image.save(
            image_name,
            InMemoryUploadedFile(
                im,       # file
                None,               # field_name
                image_name,           # file name
                'image/jpeg',       # content_type
                im.tell,  # size
                None)               # content_type_extra
        )

        # Save thumbnail
        imp = Image.open(image_data)
        print(imp)
        imp.thumbnail((100,100))
        im_thumb_data = io.BytesIO()
        imp.save(im_thumb_data, format='jpeg')
        im_thumb = ContentFile(im_thumb_data.getvalue())
        image_thumb_name = str(self.id) + "_thumbnail" + ".jpg"
        self.schedule_image_thumb.save(
            image_thumb_name,
            InMemoryUploadedFile(
                im_thumb,       # file
                None,               # field_name
                image_thumb_name,           # file name
                'image/jpeg',       # content_type
                im_thumb.tell,  # size
                None)               # content_type_extra
        )

        # self.save()

    def day_today(self):
        return dt.date.today().strftime("%A")
    
    def subjects_today(self):
        today = dt.date.today()
        days_elapsed = (today - self.start_date).days
        day = days_elapsed % self.period

        subj_today = [None for _ in range(self.time_schema.entries())]
        schema = self.time_schema.get_pythonized_items()
        for subject in self.subjects.all():
            # days_and_orders
            dao = json.loads(subject.days_and_orders)
            orders = dao.get(str(day))
            if orders != None:
                for order in orders:
                    with suppress(IndexError):
                        start = dt.time.fromisoformat(schema[order]['start'])
                        stop = dt.time.fromisoformat(schema[order]['stop'])
                        subject.start = start.strftime("%H:%M")
                        subject.stop = stop.strftime("%H:%M")
                        subj_today[order] = subject
        return subj_today


class Subject(models.Model):
    room = models.ForeignKey(Room, on_delete=models.CASCADE,
                             related_name='subjects')
    # owner = models.ForeignKey(settings.AUTH_USER_MODEL,
    #                           on_delete=models.CASCADE)
    name = models.CharField(max_length=200)
    days_and_orders = models.TextField(default="[]")
    lector = models.CharField(max_length=200, blank=True, default='')
    extra = models.TextField(blank=True, default='')

    def __str__(self):
        return f"{self.name}"
