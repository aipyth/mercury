# Generated by Django 3.1.1 on 2021-02-28 21:07

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('schedule', '0005_auto_20210228_2053'),
    ]

    operations = [
        migrations.AlterField(
            model_name='room',
            name='slug',
            field=models.SlugField(null=True, unique=True),
        ),
    ]

