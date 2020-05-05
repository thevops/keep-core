import React from "react"
import { Route, Switch, Redirect } from "react-router-dom"
import { withContractsDataContext } from "./ContractsDataContextProvider"
import Loadable from "./Loadable"
import { NotFound404 } from "./NotFound404"
import withWeb3Context from "./WithWeb3Context"
import OperatorPage from "../pages/OperatorPage"
import AuthorizerPage from "../pages/AuthorizerPage"
import RewardsPage from "../pages/RewardsPage"
import CreateTokenGrantPage from "../pages/CreateTokenGrantPage"
import TokenGrantsPage from "../pages/TokenGrantsPage"
import TokensPageContainer from "../pages/TokensPageContainer"
import ChooseWallet from "./ChooseWallet"

class Routing extends React.Component {
  renderContent() {
    const {
      isKeepTokenContractDeployer,
      contractsDataIsFetching,
      web3: { error, provider },
    } = this.props

    if (!provider) {
      return <ChooseWallet />
    }

    if (error) {
      return null
    }

    return contractsDataIsFetching ? (
      <Loadable />
    ) : (
      <Switch>
        <Route path="/tokens" component={TokensPageContainer} />
        <Route exact path="/operations" component={OperatorPage} />
        <Route exact path="/rewards" component={RewardsPage} />
        <Route exact path="/authorizer" component={AuthorizerPage} />
        <Route exact path="/token-grants" component={TokenGrantsPage} />
        {isKeepTokenContractDeployer && (
          <Route
            exact
            path="/create-token-grants"
            component={CreateTokenGrantPage}
          />
        )}
        <Route exact path="/">
          <Redirect to="/tokens" />
        </Route>
        <Route path="*">
          <NotFound404 />
        </Route>
      </Switch>
    )
  }

  render() {
    return <>{this.renderContent()}</>
  }
}

export default withWeb3Context(withContractsDataContext(Routing))
