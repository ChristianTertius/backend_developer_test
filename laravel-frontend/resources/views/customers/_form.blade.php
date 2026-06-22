@php
    $oldFamilies = old('families', isset($customer) ? $customer->families->toArray() : []);
@endphp
<div class="card">
    <div class="card-body">
        <div class="row g-3">
            <div class="col-md-6">
                <label class="form-label">Nama</label>
                <input type="text" name="cst_name" class="form-control" placeholder="Masukan nama anda"
                       value="{{ old('cst_name', isset($customer) ? trim($customer->cst_name) : '') }}" required>
            </div>
            <div class="col-md-3">
                <label class="form-label">Tanggal Lahir</label>
                <input type="date" name="cst_dob" class="form-control"
                       value="{{ old('cst_dob', isset($customer) ? $customer->cst_dob : '') }}" required>
            </div>
            <div class="col-md-3">
                <label class="form-label">Kewarganegaraan</label>
                <select name="nationality_id" class="form-select" required>
                    <option value="">Pilih kewarganegaraan</option>
                    @foreach ($nationalities as $n)
                        <option value="{{ $n->nationality_id }}"
                            {{ (string) old('nationality_id', isset($customer) ? $customer->nationality_id : '') === (string) $n->nationality_id ? 'selected' : '' }}>
                            {{ $n->nationality_name }}
                        </option>
                    @endforeach
                </select>
            </div>
            <div class="col-md-6">
                <label class="form-label">No. Telepon</label>
                <input type="text" name="cst_phoneNum" class="form-control"
                       value="{{ old('cst_phoneNum', isset($customer) ? $customer->cst_phoneNum : '') }}" required>
            </div>
            <div class="col-md-6">
                <label class="form-label">Email</label>
                <input type="email" name="cst_email" class="form-control"
                       value="{{ old('cst_email', isset($customer) ? $customer->cst_email : '') }}" required>
            </div>
        </div>
        <hr class="my-4">
        <div class="d-flex justify-content-between align-items-center mb-2">
            <h5 class="mb-0">Keluarga</h5>
            <button type="button" id="add-family" class="btn btn-link p-0">+ Tambah Keluarga</button>
        </div>
        <div id="families-wrapper">
            @foreach ($oldFamilies as $i => $fam)
                <div class="row g-2 align-items-end mb-2 family-row">
                    <div class="col-md-4">
                        <label class="form-label">Nama</label>
                        <input type="text" name="families[{{ $i }}][fl_name]" class="form-control" placeholder="Masukan Nama" value="{{ $fam['fl_name'] ?? '' }}">
                    </div>
                    <div class="col-md-3">
                        <label class="form-label">Relasi</label>
                        <input type="text" name="families[{{ $i }}][fl_relation]" class="form-control" placeholder="mis. Anak / Istri" value="{{ $fam['fl_relation'] ?? '' }}">
                    </div>
                    <div class="col-md-3">
                        <label class="form-label">Tanggal Lahir</label>
                        <input type="date" name="families[{{ $i }}][fl_dob]" class="form-control" value="{{ $fam['fl_dob'] ?? '' }}">
                    </div>
                    <div class="col-md-2">
                        <button type="button" class="btn btn-danger w-100 remove-family">Hapus</button>
                    </div>
                </div>
            @endforeach
        </div>
        <div class="mt-4">
            <button type="submit" class="btn btn-success">Simpan</button>
            <a href="{{ route('customers.index') }}" class="btn btn-secondary">Batal</a>
        </div>
    </div>
</div>
@push('scripts')
<script>
    (function () {
        const wrapper = document.getElementById('families-wrapper');
        let index = wrapper.querySelectorAll('.family-row').length;

        function rowTemplate(i) {
            return `
            <div class="row g-2 align-items-end mb-2 family-row">
                <div class="col-md-4">
                    <label class="form-label">Nama</label>
                    <input type="text" name="families[${i}][fl_name]" class="form-control" placeholder="Masukan Nama">
                </div>
                <div class="col-md-3">
                    <label class="form-label">Relasi</label>
                    <input type="text" name="families[${i}][fl_relation]" class="form-control" placeholder="mis. Anak / Istri">
                </div>
                <div class="col-md-3">
                    <label class="form-label">Tanggal Lahir</label>
                    <input type="date" name="families[${i}][fl_dob]" class="form-control">
                </div>
                <div class="col-md-2">
                    <button type="button" class="btn btn-danger w-100 remove-family">Hapus</button>
                </div>
            </div>`;
        }

        document.getElementById('add-family').addEventListener('click', function () {
            wrapper.insertAdjacentHTML('beforeend', rowTemplate(index));
            index++;
        });

        wrapper.addEventListener('click', function (e) {
            if (e.target.classList.contains('remove-family')) {
                e.target.closest('.family-row').remove();
            }
        });
    })();
</script>
@endpush
