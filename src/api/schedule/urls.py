from django.urls import path, include
from rest_framework import routers
from rest_framework.schemas import get_schema_view
from django.views.generic import TemplateView
from rest_framework.authtoken.views import obtain_auth_token

from schedule.views import (TimeSchemaViewSet, RoomViewSet, SubjectsViewSet,
                            CustomUserCreate, RoomDetailView)

router = routers.DefaultRouter()
router.register(r'timeschemes', TimeSchemaViewSet, basename='timeschema')
router.register(r'rooms', RoomViewSet, basename='room')
router.register(r'subjects', SubjectsViewSet, basename='subject')


# urlpatterns = router.urls
urlpatterns = [
    path('api/', include(router.urls)),
    path('openapi', get_schema_view(
        title="Mercury",
        description="Mercury API Shema",
        version="1.0.0"
    ), name='openapi-schema'),
    path('redoc/', TemplateView.as_view(
        template_name='redoc.html',
        extra_context={'schema_url': 'openapi-schema'}
    ), name='redoc'),

    path('api/create-user/', CustomUserCreate.as_view()),
    path('api/api-token-auth/', obtain_auth_token, name='api_token_auth'),
    path('', TemplateView.as_view(template_name='index.html'), name='index'),
    path('room/<int:pk>/', RoomDetailView.as_view()),
    path('room/<slug:slug>/', RoomDetailView.as_view()),
    path('<slug:slug>/', TemplateView.as_view(template_name='index.html'))
]
