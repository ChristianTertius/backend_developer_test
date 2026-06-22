<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Nationality extends Model
{
    protected $table = 'nationality';
    protected $primaryKey = 'nationality_id';
    public $timestamps = false;

    protected $fillable = [
        'nationality_name',
        'nationality_code',
    ];

    public function customers(): HasMany
    {
        return $this->hasMany(Customer::class, 'nationality_id', 'nationality_id');
    }
}
