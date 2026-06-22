<?php

namespace App\Http\Controllers;

use App\Models\Customer;
use App\Models\Nationality;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class CustomerController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $customers = Customer::with(['nationality', 'families'])->orderBy('cst_id')->get();

        return view('customers.index', compact('customers'));
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        $nationalities = Nationality::orderBy('nationality_name')->get();
        $customer = new Customer();

        return view('customers.create', compact('nationalities', 'customer'));
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $data = $this->validateData($request);

        DB::transaction(function () use ($data) {
            $customer = Customer::create([
                'nationality_id' => $data['nationality_id'],
                'cst_name'       => $data['cst_name'],
                'cst_dob'        => $data['cst_dob'],
                'cst_phoneNum'   => $data['cst_phoneNum'],
                'cst_email'      => $data['cst_email'],
            ]);

            $this->syncFamilies($customer, $data['families'] ?? []);
        });

        return redirect()->route('customers.index')
            ->with('success', 'Data customer berhasil ditambahkan.');
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Customer  $customer
     * @return \Illuminate\Http\Response
     */
    public function show(Customer $customer)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Customer  $customer
     * @return \Illuminate\Http\Response
     */
    public function edit(Customer $customer)
    {
        $customer->load('families');
        $nationalities = Nationality::orderBy('nationality_name')->get();

        return view('customers.edit', compact('customer', 'nationalities'));
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Customer  $customer
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Customer $customer)
    {
        $data = $this->validateData($request);

        DB::transaction(function () use ($customer, $data) {
            $customer->update([
                'nationality_id' => $data['nationality_id'],
                'cst_name'       => $data['cst_name'],
                'cst_dob'        => $data['cst_dob'],
                'cst_phoneNum'   => $data['cst_phoneNum'],
                'cst_email'      => $data['cst_email'],
            ]);

            // Reset keluarga lalu insert ulang sesuai input form.
            $customer->families()->delete();
            $this->syncFamilies($customer, $data['families'] ?? []);
        });

        return redirect()->route('customers.index')
            ->with('success', 'Data customer berhasil diperbarui.');
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Customer  $customer
     * @return \Illuminate\Http\Response
     */
    public function destroy(Customer $customer)
    {
        $customer->delete();

        return redirect()->route('customers.index')
            ->with('success', 'Data customer berhasil dihapus.');
    }

    private function validateData(Request $request): array
    {
        return $request->validate([
            'nationality_id'      => ['required', 'exists:nationality,nationality_id'],
            'cst_name'            => ['required', 'string', 'max:50'],
            'cst_dob'             => ['required', 'date'],
            'cst_phoneNum'        => ['required', 'string', 'max:20'],
            'cst_email'           => ['required', 'email', 'max:50'],
            'families'            => ['array'],
            'families.*.fl_relation' => ['required_with:families.*.fl_name', 'nullable', 'string', 'max:50'],
            'families.*.fl_name'     => ['required_with:families.*.fl_relation', 'nullable', 'string', 'max:50'],
            'families.*.fl_dob'      => ['nullable', 'date'],
        ]);
    }


    private function syncFamilies(Customer $customer, array $families): void
    {
        foreach ($families as $f) {
            if (empty($f['fl_name']) && empty($f['fl_relation'])) {
                continue;
            }
            $customer->families()->create([
                'fl_relation' => $f['fl_relation'] ?? '',
                'fl_name'     => $f['fl_name'] ?? '',
                'fl_dob'      => $f['fl_dob'] ?? '',
            ]);
        }
    }
}
