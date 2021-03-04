from rest_framework import generics
from rest_framework import permissions
from rest_framework import viewsets
from rest_framework.response import Response

from users.models import CustomUser

from schedule.models import TimeSchema, Room, Subject
from schedule.serializers import (TimeSchemaSerializer, RoomSerializer,
                                  SubjectSerializer, CustomUserSerializer)
from schedule.permissions import IsOwnerOrReadOnly, IsOwner

from django.views.generic import DetailView


class TimeSchemaViewSet(viewsets.ModelViewSet):
    serializer_class = TimeSchemaSerializer
    queryset = TimeSchema.objects.all()
    permission_classes = [
        permissions.IsAuthenticatedOrReadOnly,
        IsOwnerOrReadOnly,
    ]

    def get_queryset(self):
        queryset = TimeSchema.objects.all()
        search = self.request.query_params.get('search', None)
        if search is not None:
            queryset = queryset.filter(name__startswith=search)
        return queryset


class CustomUserCreate(generics.CreateAPIView):
    queryset = CustomUser.objects.all()
    serializer_class = CustomUserSerializer


class RoomViewSet(viewsets.ModelViewSet):
    """
    ViewSet for viewing rooms associated with the user
    """
    serializer_class = RoomSerializer
    permission_classes = [
        permissions.IsAuthenticatedOrReadOnly,
        IsOwnerOrReadOnly,
    ]

    def get_queryset(self):
        if self.request.user.is_anonymous:
            queryset = Room.objects.filter(public=True)
        else:
            queryset = self.request.user.room_set.all()
        return queryset

    def perform_create(self, serializer):
        instance = serializer.save(owner=self.request.user)
        instance.generate_schedule_image()
        instance.save()
        # print(dir(serializer))

    def retrieve(self, request, pk=None):
        room = Room.objects.get(pk=pk)
        if room.schedule_image == None:
            room.generate_schedule_image()
        return super().retrieve(request, pk)

    def update(self, request, pk=None):
        super().update(request, pk)
        room = Room.objects.get(pk=pk)
        room.generate_schedule_image()
        serializer = self.get_serializer(room)
        # serializer.is_valid(raise_exception=True)
        return Response(serializer.data)



class SubjectsViewSet(viewsets.ModelViewSet):
    queryset = Subject.objects.all()
    serializer_class = SubjectSerializer
    permission_classes = [permissions.IsAuthenticatedOrReadOnly,
                          IsOwnerOrReadOnly, ]


class RoomDetailView(DetailView):
    model = Room
