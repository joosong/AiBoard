package com.aiboard.app

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.unit.dp

data class AccountConfigUi(
    val provider: String,
    val remark: String,
    val apiKeyMask: String
)

data class SnapshotUi(
    val provider: String,
    val remark: String,
    val status: String,
    val primary: String,
    val secondary: String,
    val reset: String,
    val accent: Color
)

object DemoSnapshotRepository {
    val snapshots = listOf(
        SnapshotUi("OpenAI", "公司组织", "OK", "5h: 21k tokens", "Week: admin usage enabled", "Reset 15:00", Color(0xFF166B5B)),
        SnapshotUi("MiniMax", "项目备用", "Local", "5h window active", "Week quota local fallback", "Reset 15:00", Color(0xFF9B5A18)),
        SnapshotUi("DeepSeek", "个人余额", "OK", "CNY 12.50", "Granted 2.50 / Top-up 10.00", "Live balance", Color(0xFF3442A8))
    )

    val configs = listOf(
        AccountConfigUi("OpenAI", "公司组织", "sk-...A9f2"),
        AccountConfigUi("MiniMax", "项目备用", "mk-...0C21"),
        AccountConfigUi("DeepSeek", "个人余额", "ds-...91bA")
    )
}

@Composable
fun AiBoardApp(
    posture: DevicePosture,
    snapshots: List<SnapshotUi>,
    configs: List<AccountConfigUi>
) {
    MaterialTheme {
        Surface(
            modifier = Modifier
                .fillMaxSize()
                .background(Color(0xFFF7F7F4)),
            color = Color(0xFFF7F7F4)
        ) {
            if (posture == DevicePosture.Landscape) {
                LandscapeDashboard(snapshots)
            } else {
                PortraitDashboard(snapshots, configs)
            }
        }
    }
}

@Composable
private fun PortraitDashboard(snapshots: List<SnapshotUi>, configs: List<AccountConfigUi>) {
    var selected by remember { mutableIntStateOf(0) }
    val current = snapshots[selected.coerceIn(0, snapshots.lastIndex)]

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(18.dp),
        verticalArrangement = Arrangement.spacedBy(16.dp)
    ) {
        Header("AI API 额度看板", "${configs.size} 个配置")
        SnapshotBanner(current, Modifier.fillMaxWidth())
        Row(horizontalArrangement = Arrangement.spacedBy(10.dp)) {
            Button(onClick = { selected = (selected + snapshots.size - 1) % snapshots.size }) { Text("上一张") }
            Button(onClick = { selected = (selected + 1) % snapshots.size }) { Text("下一张") }
            Button(onClick = {}) { Text("刷新") }
        }
        SettingsPanel(configs)
    }
}

@Composable
private fun LandscapeDashboard(snapshots: List<SnapshotUi>) {
    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(18.dp),
        verticalArrangement = Arrangement.spacedBy(14.dp)
    ) {
        Header("横屏轮播", "重力感应已启用")
        LazyRow(horizontalArrangement = Arrangement.spacedBy(14.dp)) {
            items(snapshots) { snapshot ->
                SnapshotBanner(snapshot, Modifier.width(360.dp))
            }
        }
    }
}

@Composable
private fun Header(title: String, subtitle: String) {
    Row(
        modifier = Modifier.fillMaxWidth(),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.SpaceBetween
    ) {
        Text(title, style = MaterialTheme.typography.headlineSmall, fontWeight = FontWeight.Bold)
        Text(subtitle, style = MaterialTheme.typography.bodyMedium, color = Color(0xFF5F6864))
    }
}

@Composable
private fun SnapshotBanner(snapshot: SnapshotUi, modifier: Modifier = Modifier) {
    Card(
        modifier = modifier.height(210.dp),
        shape = RoundedCornerShape(8.dp),
        colors = CardDefaults.cardColors(containerColor = snapshot.accent)
    ) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(18.dp),
            verticalArrangement = Arrangement.SpaceBetween
        ) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Text(snapshot.provider, color = Color.White, style = MaterialTheme.typography.titleLarge, fontWeight = FontWeight.Bold)
                Text(snapshot.status, color = Color.White)
            }
            Column {
                Text(snapshot.primary, color = Color.White, style = MaterialTheme.typography.headlineSmall, fontWeight = FontWeight.Bold)
                Spacer(Modifier.height(6.dp))
                Text(snapshot.secondary, color = Color(0xFFEAF4EF), style = MaterialTheme.typography.bodyMedium)
            }
            Text(snapshot.reset, color = Color(0xFFEAF4EF), style = MaterialTheme.typography.bodySmall)
        }
    }
}

@Composable
private fun SettingsPanel(configs: List<AccountConfigUi>) {
    Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Text("设置配置", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
        configs.forEach { config ->
            Card(shape = RoundedCornerShape(8.dp), colors = CardDefaults.cardColors(containerColor = Color.White)) {
                Row(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(14.dp),
                    horizontalArrangement = Arrangement.SpaceBetween
                ) {
                    Column {
                        Text(config.remark, fontWeight = FontWeight.SemiBold)
                        Text(config.provider, color = Color(0xFF5F6864))
                    }
                    Text(config.apiKeyMask, color = Color(0xFF5F6864))
                }
            }
        }
        Box(Modifier.fillMaxWidth()) {
            OutlinedTextField(
                value = "",
                onValueChange = {},
                modifier = Modifier.fillMaxWidth(),
                label = { Text("新增 API Key") },
                visualTransformation = PasswordVisualTransformation()
            )
        }
    }
}
