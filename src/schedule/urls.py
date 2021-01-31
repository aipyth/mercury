from django.urls import path, include
from rest_framework import routers
from rest_framework.schemas import get_schema_view
from django.views.generic import TemplateView

from schedule.views import (TimeShemeViewSet, RoomViewSet, SubjectsViewSet,
                            CustomUserCreate)

router = routers.DefaultRouter()
router.register(r'timeshemes', TimeShemeViewSet, basename='timesheme')
router.register(r'rooms', RoomViewSet, basename='room')
router.register(r'subjects', SubjectsViewSet, basename='subject')


# urlpatterns = router.urls
urlpatterns = [
    path('/api/', include(router.urls)),
    path('openapi', get_schema_view(
        title="Mercury",
        description="Mercury API Shema",
        version="1.0.0"
    ), name='openapi-schema'),
    path('redoc/', TemplateView.as_view(
        template_name='redoc.html',
        extra_context={'schema_url': 'openapi-schema'}
    ), name='redoc'),

    path('create-user/', CustomUserCreate.as_view()),
    path('', TemplateView.as_view(template_name='index.html'), name='index')
]
