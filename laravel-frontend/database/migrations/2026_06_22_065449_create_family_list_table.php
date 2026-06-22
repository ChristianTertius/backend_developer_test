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
        Schema::create('family_list', function (Blueprint $table) {
            $table->id('fl_id');
            $table->unsignedBigInteger('cst_id');
            $table->string('fl_relation', 50);
            $table->string('fl_name', 50);
            $table->string('fl_dob', 50);

            $table->foreign('cst_id')
                ->references('cst_id')
                ->on('customer')
                ->cascadeOnDelete();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('family_list');
    }
};
