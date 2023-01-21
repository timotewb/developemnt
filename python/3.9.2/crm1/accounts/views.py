from django.shortcuts import render
from django.http import HttpResponse
# Create your views here.

def home(requese):
    return HttpResponse('Home Page')

def products(response):
    return HttpResponse('Products Page')

def customer(response):
    return HttpResponse('Customer Page')