# Generated by Django 3.1.1 on 2021-02-28 20:53

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('schedule', '0004_auto_20210228_1306'),
    ]

    operations = [
        migrations.AlterField(
            model_name='room',
            name='schedule_image',
            field=models.ImageField(blank=True, default='', upload_to='schedules'),
        ),
        migrations.AlterField(
            model_name='room',
            name='schedule_image_thumb',
            field=models.ImageField(blank=True, default='', upload_to='schedules'),
        ),
        migrations.AlterField(
            model_name='subject',
            name='extra',
            field=models.TextField(blank=True, default=''),
        ),
        migrations.AlterField(
            model_name='subject',
            name='lector',
            field=models.CharField(blank=True, default='', max_length=200),
        ),
    ]
