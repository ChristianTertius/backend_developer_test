<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Customer extends Model
{
    protected $table = 'customer';
    protected $primaryKey = 'cst_id';
    public $timestamps = false;

    protected $fillable = [
        'nationality_id',
        'cst_name',
        'cst_dob',
        'cst_phoneNum',
        'cst_email',
    ];

    public function nationality(): BelongsTo
    {
        return $this->belongsTo(Nationality::class, 'nationality_id', 'nationality_id');
    }

    public function families(): HasMany
    {
        return $this->hasMany(Family::class, 'cst_id', 'cst_id');
    }
}
