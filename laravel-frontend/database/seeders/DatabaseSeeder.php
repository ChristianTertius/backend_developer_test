<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;

use App\Models\Nationality;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        $data = [
            ['nationality_name' => 'Indonesia',     'nationality_code' => 'ID'],
            ['nationality_name' => 'Singapore',     'nationality_code' => 'SG'],
            ['nationality_name' => 'Malaysia',      'nationality_code' => 'MY'],
            ['nationality_name' => 'United States', 'nationality_code' => 'US'],
            ['nationality_name' => 'Japan',         'nationality_code' => 'JP'],
        ];
        foreach ($data as $row) {
            Nationality::firstOrCreate(
                ['nationality_code' => $row['nationality_code']],
                $row
            );
        }
    }
}
