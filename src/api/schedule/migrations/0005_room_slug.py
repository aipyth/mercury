# Generated by Django 3.1.1 on 2021-02-22 14:44

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('schedule', '0004_subject_owner'),
    ]

    operations = [
        migrations.AddField(
            model_name='room',
            name='slug',
            field=models.SlugField(null=True),
        ),
    ]