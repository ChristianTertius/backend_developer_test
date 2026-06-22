@extends('layouts.app')
@section('title', 'Tambah User')
@section('content')
<h3 class="mb-3">Tambah User</h3>
<form action="{{ route('customers.store') }}" method="POST">
    @csrf
    @include('customers._form')
</form>
@endsection
