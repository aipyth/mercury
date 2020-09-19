from django.contrib.auth.models import User
from rest_framework import generics
from rest_framework import permissions
from rest_framework import viewsets

from schedule.models import Room, Subject
from schedule.serializers import RoomSerializer, UserSerializer, SubjectSerializer
from schedule.permissions import IsOwnerOrReadOnly, IsOwner


class UserCreate(generics.CreateAPIView):
    queryset = User.objects.all()
    serializer_class = UserSerializer


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
    permission_classes = [IsOwner]
