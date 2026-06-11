package com.aiboard.app

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import kotlin.math.abs

enum class DevicePosture {
    Portrait,
    Landscape
}

class SensorController(
    context: Context,
    private val onPostureChanged: (DevicePosture) -> Unit
) : SensorEventListener {
    private val sensorManager = context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
    private val accelerometer = sensorManager.getDefaultSensor(Sensor.TYPE_ACCELEROMETER)
    private var lastPosture = DevicePosture.Portrait

    fun start() {
        accelerometer?.let {
            sensorManager.registerListener(this, it, SensorManager.SENSOR_DELAY_UI)
        }
    }

    fun stop() {
        sensorManager.unregisterListener(this)
    }

    override fun onSensorChanged(event: SensorEvent) {
        val x = event.values[0]
        val y = event.values[1]
        val next = if (abs(x) > abs(y) + 1.5f) DevicePosture.Landscape else DevicePosture.Portrait
        if (next != lastPosture) {
            lastPosture = next
            onPostureChanged(next)
        }
    }

    override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) = Unit
}
