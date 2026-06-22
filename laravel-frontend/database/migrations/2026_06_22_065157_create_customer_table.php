<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('customer', function (Blueprint $table) {
            $table->id('cst_id');
            $table->unsignedBigInteger('nationality_id');
            $table->char('cst_name', 50);
            $table->date('cst_dob');
            $table->string('cst_phoneNum', 20);
            $table->string('cst_email', 50);

            $table->foreign('nationality_id')
                ->references('nationality_id')
                ->on('nationality');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('customer');
    }
};
