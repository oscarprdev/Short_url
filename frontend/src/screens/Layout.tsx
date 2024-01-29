import { Route } from 'wouter';
import HomeScreen from './HomeScreen';
import Footer from '../components/Footer';
import AuthScreen from './AuthScreen';
import Header from '../components/Header';
import Background from '../components/Background';
import UserScreenContainer from '../components/containers/UserScreenContainer';
import { Toaster } from '../components/ui/toaster';
import RedirectScreen from './RedirectScreen';

const Layout = () => {
	return (
		<>
			<Header />
			<Route path='/user/:id'>
				{(params) => (
					<AuthScreen userId={params.id}>
						<UserScreenContainer />
					</AuthScreen>
				)}
			</Route>
			<Route path='/'>
				<HomeScreen />
			</Route>
			<Route path='/:id'>{(params) => <RedirectScreen urlId={params.id} />}</Route>
			<Footer />
			<Toaster />
			<Background />
		</>
	);
};

export default Layout;
