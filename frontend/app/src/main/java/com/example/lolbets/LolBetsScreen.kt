package com.example.lolbets

import android.content.Intent
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.AccountCircle
import androidx.compose.material.icons.filled.KeyboardArrowLeft
import androidx.compose.material.icons.rounded.AddCircle
import androidx.compose.material.icons.rounded.Home
import androidx.compose.material.icons.rounded.Settings
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.lolbets.model.BottomNavItem
import com.example.lolbets.model.Game
import com.example.lolbets.model.League
import com.example.lolbets.model.Team
import com.example.lolbets.model.User
import com.example.lolbets.ui.BetScreen
import com.example.lolbets.ui.FocusedGameViewModel
import com.example.lolbets.ui.GamesScreen
import com.example.lolbets.ui.HighlightScreen
import com.example.lolbets.ui.ProfileScreen
import com.google.android.gms.auth.api.signin.GoogleSignInClient
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.lolbets.ui.GamesViewModel
import com.example.lolbets.ui.HomeScreen


enum class LolBetsScreen(){
    Highlight,
    Games,
    Profile,
    Bet,
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
internal fun LolBetsTopAppBar(onProfileClicked: () -> Unit, onArrowBackClicked: () -> Unit, modifier: Modifier = Modifier) {
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
            IconButton(onClick = { onArrowBackClicked() }) {
                Icon(
                    imageVector = Icons.Filled.KeyboardArrowLeft,//Icons.Filled.ArrowBack
                    contentDescription = "Localized description"
                )
            }
        },
        actions = {
            Text(text = "10$")
            IconButton(onClick = { onProfileClicked() }) {
                Icon(
                    imageVector = Icons.Filled.AccountCircle,
                    contentDescription = "Localized description"
                )
            }
        }
    )
}


@Composable
internal fun LolBetsBottomAppBar(items: List<BottomNavItem>, modifier: Modifier = Modifier) {
    // NavController is passed via parameter
    //val backStackEntry = navController.currentBackStackEntryAsState()
    var selectedItem by remember { mutableStateOf(0) }

    NavigationBar(
        //containerColor = MaterialTheme.colors.primary,
    ) {
        items.forEachIndexed { index, item ->
            NavigationBarItem(
                icon = { Icon(imageVector = item.icon, contentDescription = "${item.name} Icon") },
                label = { Text(item.name, fontWeight = FontWeight.SemiBold) },
                selected = selectedItem == index,
                //onClick = { selectedItem = index },
                onClick = {
                    selectedItem = index
                    item.onButtonClicked()
                },

            )
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LolBetsApp2(
    mGoogleSignInClient: GoogleSignInClient,
    modifier: Modifier = Modifier,
    viewModel: FocusedGameViewModel = viewModel(),
    navController: NavHostController = rememberNavController(),
) {
    val items = listOf(
        BottomNavItem(
            name = "Home",
            onButtonClicked = { navController.navigate(LolBetsScreen.Games.name) },
            icon = Icons.Rounded.Home,
        ),
        BottomNavItem(
            name = "Create",
            onButtonClicked = { navController.navigate(LolBetsScreen.Highlight.name) },
            icon = Icons.Rounded.AddCircle,
        ),
        BottomNavItem(
            name = "Settings",
            onButtonClicked = { navController.navigate(LolBetsScreen.Profile.name) },
            icon = Icons.Rounded.Settings,
        ),
    )

    Scaffold(
        topBar = {
            LolBetsTopAppBar( onProfileClicked = { navController.navigate(LolBetsScreen.Profile.name) }, onArrowBackClicked = { navController.navigateUp() }, modifier)
        },
        bottomBar = {
            LolBetsBottomAppBar(items, modifier)
        }
    ) { innerPadding ->
        val uiState by viewModel.uiState.collectAsState()

        NavHost(
            navController = navController,
            startDestination = LolBetsScreen.Games.name,
            modifier = Modifier.padding(innerPadding)
        ) {
            composable(route = LolBetsScreen.Games.name) {
                GamesScreen(
                    contentPadding = innerPadding,
                    onGameClicked = {
                        viewModel.setGame(it)
                        navController.navigate(LolBetsScreen.Bet.name) },
                )
            }
            composable(route = LolBetsScreen.Highlight.name) {
                HighlightScreen(
                    contentPadding = innerPadding,
                )
            }
            composable(route = LolBetsScreen.Profile.name) {
                ProfileScreen(
                    User(R.string.team_name_koi, R.drawable.koi, 10),
                    contentPadding = innerPadding,
                )
            }
            composable(route = LolBetsScreen.Bet.name) {
                BetScreen(
                    betState = uiState
                )
            }
        }


        /*val startForResult =
            rememberLauncherForActivityResult(ActivityResultContracts.StartActivityForResult()) { result: ActivityResult ->
                if (result.resultCode == Activity.RESULT_OK) {
                    val intent = result.data
                    if (result.data != null) {
                        val task: Task<GoogleSignInAccount> =
                            GoogleSignIn.getSignedInAccountFromIntent(intent)
                        task.result
                        handleSignInResult(task)
                    }
                }
            }*/

        /*Button(
            /*onClick = {
                val signInIntent: Intent = mGoogleSignInClient.getSignInIntent()
                startActivityForResult(signInIntent, RC_SIGN_IN)
            },*/
            onClick = {},
            modifier = Modifier
                .fillMaxWidth()
                .padding(innerPadding),
            shape = RoundedCornerShape(6.dp),
        ) {
            Text(text = "Sign in with Google", modifier = Modifier.padding(6.dp))
        }*/

        //LoginScreen(clientIdtest, innerPadding)
        //BetScreen(innerPadding)
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LolBetsApp(
    mGoogleSignInClient: GoogleSignInClient,
    modifier: Modifier = Modifier,
    //viewModel: GamesViewModel = viewModel()
) {
    val viewModel: GamesViewModel = viewModel()
    HomeScreen(gameUiState = viewModel.gameUiState)
}
