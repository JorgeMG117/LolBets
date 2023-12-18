package com.example.lolbets

import android.app.Activity.RESULT_OK
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.IntentSenderRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.AccountCircle
import androidx.compose.material.icons.filled.KeyboardArrowLeft
import androidx.compose.material.icons.rounded.AddCircle
import androidx.compose.material.icons.rounded.Home
import androidx.compose.material.icons.rounded.Settings
import androidx.compose.material.icons.rounded.Star
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.navigation.NavHostController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.lolbets.model.BottomNavItem
import com.example.lolbets.model.User
import com.example.lolbets.ui.BetScreen
import com.example.lolbets.viewmodel.FocusedGameViewModel
import com.example.lolbets.ui.HighlightScreen
import com.example.lolbets.ui.ProfileScreen
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.lolbets.viewmodel.GamesViewModel
import com.example.lolbets.ui.HomeScreen
import com.example.lolbets.viewmodel.UserViewModel
import androidx.compose.ui.platform.LocalLifecycleOwner
import com.example.lolbets.ui.BetsSummary
import com.example.lolbets.ui.sign_in.SignInViewModel
import com.example.lolbets.viewmodel.ActiveBetsViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.lifecycleScope
import com.example.lolbets.ui.sign_in.GoogleAuthUiClient
import com.example.lolbets.ui.sign_in.SignInScreen
import kotlinx.coroutines.launch


enum class LolBetsScreen(){
    App,
    Highlight,
    Summary,
    Games,
    Profile,
    Bet,
    Login,
}

fun onScreenChange(focusedGameViewModel: FocusedGameViewModel) {
    focusedGameViewModel.disconnectWebSocket()
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
internal fun LolBetsTopAppBar(user : User, onProfileClicked: () -> Unit, onArrowBackClicked: () -> Unit, modifier: Modifier = Modifier) {
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
            val coinsString = user.coins.toString()
            Text(text = "$coinsString\$")
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
internal fun LolBetsBottomAppBar(items: List<BottomNavItem>, onNavigationClicked: () -> Unit, modifier: Modifier = Modifier) {
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
                    onNavigationClicked()
                    item.onButtonClicked()
                },

            )
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LolBetsApp(
    googleAuthUiClient: GoogleAuthUiClient,
    modifier: Modifier = Modifier,
    navController: NavHostController = rememberNavController(),
) {
    //Maybe: Screen view model to manage the botton and up bars?
    // Use LaunchedEffect to observe the NavController
    var shouldShowAppBars by remember { mutableStateOf(false) }

    val viewModelGames: GamesViewModel = viewModel()

    //Active Bets
    val viewModelActiveBets: ActiveBetsViewModel = viewModel {
        ActiveBetsViewModel(0)
    }

    // Focused game view model
    val viewModel: FocusedGameViewModel = viewModel {
        FocusedGameViewModel(
            {
                viewModelGames.updateGame(it)
            }
        )
    }

    //UserViewModel
    val viewModelUser : UserViewModel = viewModel()
    val userState by viewModelUser.userState.collectAsState()

    val items = listOf(
        BottomNavItem(
            name = "Games",
            onButtonClicked = { navController.navigate(LolBetsScreen.Games.name) },
            icon = Icons.Rounded.Home,
        ),
        BottomNavItem(
            name = "Bets",
            onButtonClicked = { navController.navigate(LolBetsScreen.Summary.name) },
            icon = Icons.Rounded.AddCircle,
        ),
        BottomNavItem(
            name = "Highlights",
            onButtonClicked = { navController.navigate(LolBetsScreen.Highlight.name) },
            icon = Icons.Rounded.Star,
        ),
    )

    Scaffold(
        topBar = {
            if (shouldShowAppBars) {
                LolBetsTopAppBar(
                    userState,
                    onProfileClicked = { navController.navigate(LolBetsScreen.Profile.name) },
                    onArrowBackClicked = {
                        onScreenChange(viewModel)
                        navController.navigateUp()
                    },
                    modifier
                )
            }
        },
        bottomBar = {
            if (shouldShowAppBars) {
                /*val navBackStackEntry by navController.currentBackStackEntryAsState()
            val sceen = navBackStackEntry!!.id
            println("fdasjfkl")
            println(sceen)

            if (sceen != LolBetsScreen.Login.name) {
                LolBetsBottomAppBar(items, modifier)
            }*/
                LolBetsBottomAppBar(items, {onScreenChange(viewModel)}, modifier)
            }
        }
    ) { innerPadding ->
        val uiState by viewModel.uiState.collectAsState()

        NavHost(
            navController = navController,
            //startDestination = LolBetsScreen.Games.name,
            startDestination = LolBetsScreen.Login.name,
            modifier = Modifier.padding(innerPadding)
        ) {
            composable(route = LolBetsScreen.Login.name) {
                //navController.navigate(LolBetsScreen.Games.name)
                val viewModelLogin = viewModel<SignInViewModel>()
                val loginState by viewModelLogin.state.collectAsStateWithLifecycle()

                //Check if user is already logged in
                LaunchedEffect(key1 = Unit) {
                    if(googleAuthUiClient.getSignedInUser() != null) {
                        //Cojer tambien info del usuario del backend
                        viewModelUser.getUserInfo(
                            googleAuthUiClient.getSignedInUser()
                        ) {
                            viewModelActiveBets.updateActiveBets(it)
                            viewModel.setUserId(it)
                        }

                        shouldShowAppBars = true

                        navController.navigate(LolBetsScreen.Games.name)
                    }
                }

                val lifecycleOwner = LocalLifecycleOwner.current
                val launcher = rememberLauncherForActivityResult(
                    contract = ActivityResultContracts.StartIntentSenderForResult(),
                    onResult = { result ->
                        if(result.resultCode == RESULT_OK) {
                            lifecycleOwner.lifecycleScope.launch {
                                val signInResult = googleAuthUiClient.signInWithIntent(
                                    intent = result.data ?: return@launch
                                )
                                viewModelLogin.onSignInResult(signInResult)
                            }
                        }
                    }
                )

                LaunchedEffect(key1 = loginState.isSignInSuccessful) {
                    if(loginState.isSignInSuccessful) {
                        /*Toast.makeText(
                            applicationContext,
                            "Sign in successful",
                            Toast.LENGTH_LONG
                        ).show()*/
                        println("Login Success")
                        //Hablar con backend
                        //enviar que ha avido login e enviar el id, esto nos devuelve el id del usuario y sus coins
                        viewModelUser.getUserInfo(
                            googleAuthUiClient.getSignedInUser()
                        ) {
                            viewModelActiveBets.updateActiveBets(it)
                            viewModel.setUserId(it)
                        }

                        shouldShowAppBars = true

                        //Update active bets with the new user
                        viewModelActiveBets.updateActiveBets(userState.id)

                        navController.navigate(LolBetsScreen.Games.name)
                        viewModelLogin.resetState()
                    }
                }


                SignInScreen(
                    state = loginState,
                    onSignInClick = {
                        lifecycleOwner.lifecycleScope.launch {
                            val signInIntentSender = googleAuthUiClient.signIn()
                            launcher.launch(
                                IntentSenderRequest.Builder(
                                    signInIntentSender ?: return@launch
                                ).build()
                            )
                        }
                    }
                )
            }

            composable(route = LolBetsScreen.Games.name) {
                /*GamesScreen(
                    contentPadding = innerPadding,
                    onGameClicked = {
                        viewModel.setGame(it)
                        navController.navigate(LolBetsScreen.Bet.name) },
                )*/
                HomeScreen(
                    gameUiState = viewModelGames.gameUiState,
                    contentPadding = innerPadding,
                    onGameClicked = {
                        viewModel.setGame(it)
                        viewModel.connectWebSocket(
                            it.id,
                            onBetSuccess = {
                                viewModelActiveBets.addActiveBet(it)
                        })
                        navController.navigate(LolBetsScreen.Bet.name) },
                )
            }
            composable(route = LolBetsScreen.Highlight.name) {
                /*HighlightScreen(
                    contentPadding = innerPadding,
                )*/
                HighlightScreen(
                    gameUiState = viewModelGames.gameUiState,
                    contentPadding = innerPadding,
                    onGameClicked = {
                        viewModel.setGame(it)
                        viewModel.connectWebSocket(
                            it.id,
                            onBetSuccess = {
                                viewModelActiveBets.addActiveBet(it)
                            })
                        navController.navigate(LolBetsScreen.Bet.name) },
                )
            }
            composable(route = LolBetsScreen.Summary.name) {
                BetsSummary(
                    activeBetsUiState = viewModelActiveBets.activeBetsUiState,
                    contentPadding = innerPadding,
                )
            }
            composable(route = LolBetsScreen.Profile.name) {
                /*ProfileScreen(
                    userState,
                    contentPadding = innerPadding,
                )*/
                val lifecycleOwner = LocalLifecycleOwner.current
                ProfileScreen(
                    userData = userState,
                    onSignOut = {
                        lifecycleOwner.lifecycleScope.launch {
                            googleAuthUiClient.signOut()
                            /*Toast.makeText(
                                applicationContext,
                                "Signed out",
                                Toast.LENGTH_LONG
                            ).show()*/

                            shouldShowAppBars = false

                            navController.navigate(LolBetsScreen.Login.name)
                        }
                    }
                )
            }
            composable(route = LolBetsScreen.Bet.name) {
                BetScreen(
                    onBetPlaced = { value ->
                        //TODO Check if bet is valid
                        viewModelUser.placeBet(value.value)
                        viewModel.sendMessage(value)//Send bet to server
                        //TODO Add it to active bets, no need to do the request again
                        //viewModelActiveBets.addActiveBet(value)
                    },
                    /*onBetSuccess = {
                        viewModelActiveBets.addActiveBet(it)
                    },*/
                    betState = uiState
                )
            }
        }

    }
}

