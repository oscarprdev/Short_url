import { Route } from 'wouter';
import HomeScreen from './HomeScreen';
import Footer from '../components/Footer';
import AuthScreen from './AuthScreen';
import Header from '../components/Header';
import Background from '../components/Background';
import UserScreenContainer from '../components/containers/UserScreenContainer';

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
			<Footer />
			<Background />
		</>
	);
};

export default Layout;
