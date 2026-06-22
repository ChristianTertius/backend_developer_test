@extends('layouts.app')
@section('title', 'Edit User')
@section('content')
<h3 class="mb-3">Edit User</h3>
<form action="{{ route('customers.update', $customer) }}" method="POST">
    @csrf
    @method('PUT')
    @include('customers._form')
</form>
@endsection
