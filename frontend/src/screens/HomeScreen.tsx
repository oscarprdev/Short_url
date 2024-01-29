import { Redirect } from 'wouter';
import HomeContainer from '../components/containers/HomeContainer';
import { useGlobalStore } from '../store/globalState';

const HomeScreen = () => {
	const { user } = useGlobalStore();

	if (user) {
		return <Redirect to={`/user/${user.id}`} />;
	}

	return <HomeContainer />;
};

export default HomeScreen;
