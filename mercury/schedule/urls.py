from django.urls import path
from rest_framework import routers

from schedule.views import UserCreate, RoomViewSet, SubjectsViewSet

router = routers.DefaultRouter()
router.register(r'rooms', RoomViewSet, basename='room')
router.register(r'subjects', SubjectsViewSet, basename='subject')


urlpatterns = router.urls
urlpatterns += [
    # path('room', RoomList.as_view()),
    path('create-user', UserCreate.as_view())
]
