package com.example.lolbets

import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.AccountCircle
import androidx.compose.material.icons.filled.KeyboardArrowLeft
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextOverflow
import com.example.lolbets.data.GamesData
import com.example.lolbets.ui.GamesScreen
import com.example.lolbets.ui.components.GamesList

@OptIn(ExperimentalMaterial3Api::class)
@Composable
internal fun LolBetsTopAppBar(modifier: Modifier = Modifier) {
    TopAppBar(
        modifier = modifier,
        title = {
            Text(
                stringResource(R.string.app_name),
                maxLines = 1,
                overflow = TextOverflow.Ellipsis
            )
        },
        navigationIcon = {
            IconButton(onClick = { /* doSomething() */ }) {
                Icon(
                    imageVector = Icons.Filled.KeyboardArrowLeft,
                    contentDescription = "Localized description"
                )
            }
        },
        actions = {
            Text(text = "10$")
            IconButton(onClick = { /* doSomething() */ }) {
                Icon(
                    imageVector = Icons.Filled.AccountCircle,
                    contentDescription = "Localized description"
                )
            }
        }
    )
}


@Composable
internal fun LolBetsBottomAppBar(modifier: Modifier = Modifier) {
    NavigationBar {

    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LolBetsApp(modifier: Modifier = Modifier) {
    Scaffold(
        topBar = {
            LolBetsTopAppBar(modifier)
        },
        bottomBar = {
            LolBetsBottomAppBar(modifier)
        }
    ) { innerPadding ->
        GamesScreen(innerPadding)
    }
}