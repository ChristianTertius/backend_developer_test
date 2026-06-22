@extends('layouts.app')
@section('title', 'Daftar User')
@section('content')
<div class="d-flex justify-content-between align-items-center mb-3">
    <h3 class="mb-0">Daftar User</h3>
    <a href="{{ route('customers.create') }}" class="btn btn-primary">+ Tambah User</a>
</div>
<div class="card">
    <div class="table-responsive">
        <table class="table table-striped table-hover mb-0 align-middle">
            <thead class="table-dark">
                <tr>
                    <th>#</th>
                    <th>Nama</th>
                    <th>Tanggal Lahir</th>
                    <th>Kewarganegaraan</th>
                    <th>No. Telp</th>
                    <th>Email</th>
                    <th>Keluarga</th>
                    <th class="text-end">Aksi</th>
                </tr>
            </thead>
            <tbody>
                @forelse ($customers as $customer)
                    <tr>
                        <td>{{ $customer->cst_id }}</td>
                        <td>{{ trim($customer->cst_name) }}</td>
                        <td>{{ $customer->cst_dob }}</td>
                        <td>{{ optional($customer->nationality)->nationality_name }}</td>
                        <td>{{ $customer->cst_phoneNum }}</td>
                        <td>{{ $customer->cst_email }}</td>
                        <td>
                            @forelse ($customer->families as $fam)
                                <span class="badge bg-secondary">{{ $fam->fl_relation }}: {{ $fam->fl_name }}</span>
                            @empty
                                <span class="text-muted">-</span>
                            @endforelse
                        </td>
                        <td class="text-end">
                            <a href="{{ route('customers.edit', $customer) }}" class="btn btn-sm btn-warning">Edit</a>
                            <form action="{{ route('customers.destroy', $customer) }}" method="POST" class="d-inline" onsubmit="return confirm('Hapus data ini?')">
                                @csrf
                                @method('DELETE')
                                <button class="btn btn-sm btn-danger">Hapus</button>
                            </form>
                        </td>
                    </tr>
                @empty
                    <tr><td colspan="8" class="text-center text-muted py-4">Belum ada data.</td></tr>
                @endforelse
            </tbody>
        </table>
    </div>
</div>
@endsection
