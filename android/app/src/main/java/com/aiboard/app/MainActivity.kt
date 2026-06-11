package com.aiboard.app

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue

class MainActivity : ComponentActivity() {
    private lateinit var sensorController: SensorController
    private var posture by mutableStateOf(DevicePosture.Portrait)

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        sensorController = SensorController(this) { posture = it }

        setContent {
            AiBoardApp(
                posture = posture,
                snapshots = DemoSnapshotRepository.snapshots,
                configs = DemoSnapshotRepository.configs
            )
        }
    }

    override fun onResume() {
        super.onResume()
        sensorController.start()
    }

    override fun onPause() {
        sensorController.stop()
        super.onPause()
    }
}
