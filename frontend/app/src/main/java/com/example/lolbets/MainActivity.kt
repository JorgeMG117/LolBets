package com.example.lolbets

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn

import androidx.compose.foundation.lazy.items

import androidx.compose.material3.Card
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import com.example.lolbets.data.LeagueData
import com.example.lolbets.model.League
import com.example.lolbets.ui.theme.LolBetsTheme
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Row
import androidx.compose.ui.Alignment
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext


class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            LolBetsTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    LeaguesScreen()
                }
            }
        }
    }
}

@Composable
fun LeagueCard(league: League, modifier: Modifier = Modifier) {
    Card(modifier = modifier) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = modifier
                .background(Color(0xffb0bfd9))
                .fillMaxWidth()
        ) {
            Image(
                painter = painterResource(league.imageResourceId),
                contentDescription = stringResource(league.stringResourceId),
                modifier = Modifier,
                    //.fillMaxWidth()
                    //.height(194.dp),
                contentScale = ContentScale.Crop
            )
            Text(
                text = LocalContext.current.getString(league.stringResourceId),
                modifier = Modifier.padding(16.dp),
                style = MaterialTheme.typography.headlineSmall
            )
        }
    }
}


@Composable
fun LeaguesList(leaguesList: List<League>, modifier: Modifier = Modifier) {
    LazyColumn(modifier = modifier) {
        items(leaguesList) { league ->
            LeagueCard(
                league = league,
                modifier = Modifier.padding(8.dp)
            )
        }
    }
}

@Composable
fun LeaguesScreen(modifier: Modifier = Modifier) {
    LeaguesList(
        leaguesList = LeagueData().loadLeagues()
    )
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    LolBetsTheme {
        LeaguesScreen()
    }
}