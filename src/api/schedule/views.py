from rest_framework import generics
from rest_framework import permissions
from rest_framework import viewsets

from users.models import CustomUser

from schedule.models import TimeSheme, Room, Subject
from schedule.serializers import (TimeShemeSerializer, RoomSerializer,
                                  SubjectSerializer, CustomUserSerializer)
from schedule.permissions import IsOwnerOrReadOnly, IsOwner


class TimeShemeViewSet(viewsets.ModelViewSet):
    serializer_class = TimeShemeSerializer
    queryset = TimeSheme.objects.all()
    permission_classes = [
        permissions.IsAuthenticatedOrReadOnly,
        IsOwnerOrReadOnly,
    ]


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
        serializer.save(owner=self.request.user)


class SubjectsViewSet(viewsets.ModelViewSet):
    queryset = Subject.objects.all()
    serializer_class = SubjectSerializer
    permission_classes = [permissions.IsAuthenticatedOrReadOnly,
                          IsOwnerOrReadOnly, ]
