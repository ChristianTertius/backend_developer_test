<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Family extends Model
{
    protected $table = 'family_list';
    protected $primaryKey = 'fl_id';
    public $timestamps = false;

    protected $fillable = [
        'cst_id',
        'fl_relation',
        'fl_name',
        'fl_dob',
    ];

    public function customer(): BelongsTo
    {
        return $this->belongsTo(Customer::class, 'cst_id', 'cst_id');
    }
}
