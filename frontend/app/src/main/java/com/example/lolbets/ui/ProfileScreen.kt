package com.example.lolbets.ui

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.R
import com.example.lolbets.data.GamesData
import com.example.lolbets.model.User
import com.example.lolbets.ui.components.GamesList

@Composable
fun ProfileScreen(user: User, contentPadding: PaddingValues, modifier: Modifier = Modifier) {
    Column (
        //verticalArrangement = Arrangement.SpaceEvenly,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        val borderWidth = 1.dp
        Image(
            painter = painterResource(user.profilePictureResourceId),
            contentDescription = stringResource(user.userNameResourceId),
            contentScale = ContentScale.Crop,
            modifier = Modifier
                .size(50.dp)
                .border(
                    BorderStroke(borderWidth, Color.Black),
                    CircleShape
                )
                .padding(borderWidth)
                .clip(CircleShape)
        )

        Text(
            text = LocalContext.current.getString(user.userNameResourceId),
            fontWeight = FontWeight.Bold
        )

        Text(text = user.coins.toString() + "$")
    }
}
@Preview(showBackground = true)
@Composable
fun ProfilePreview() {
    ProfileScreen(User(R.string.team_name_koi, R.drawable.koi, 10), contentPadding = PaddingValues(10.dp))
}