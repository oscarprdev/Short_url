import { Route } from 'wouter';
import UserScreen from './UserScreen';
import Home from './Home';
import Footer from '../components/Footer';
import AuthScreen from './AuthScreen';
import ScreenWrapper from './ScreenWrapper';

const Layout = () => {
	return (
		<>
			<Route path='/home'>
				<AuthScreen>
					<ScreenWrapper>
						<UserScreen />
					</ScreenWrapper>
				</AuthScreen>
			</Route>
			<Route path='/'>
				<ScreenWrapper>
					<Home />
				</ScreenWrapper>
			</Route>
			<Footer />
		</>
	);
};

export default Layout;
