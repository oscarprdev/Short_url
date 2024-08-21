import Footer from '../components/Footer';
import Header from '../components/Header';
import UserScreenContainer from '../components/containers/UserScreenContainer';
import { Toaster } from '../components/ui/toaster';
import AuthScreen from './AuthScreen';
import HomeScreen from './HomeScreen';
import RedirectScreen from './RedirectScreen';
import { Route } from 'wouter';

const Layout = () => {
	return (
		<>
			<Header />
			<Route path="/">
				<HomeScreen />
			</Route>
			<Route path="/:url">{params => <RedirectScreen shortUrl={params.url} />}</Route>
			<Route path="/user/:id">
				{params => (
					<AuthScreen userId={params.id}>
						<UserScreenContainer />
					</AuthScreen>
				)}
			</Route>
			<Footer />
			<Toaster />
		</>
	);
};

export default Layout;
