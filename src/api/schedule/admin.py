from django.contrib import admin

from .models import Subject, Room, TimeSchema

admin.site.register(Subject)
admin.site.register(Room)
admin.site.register(TimeSchema)
