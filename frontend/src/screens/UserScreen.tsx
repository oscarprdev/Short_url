import { useUserLogged } from '../hooks/useUserLogged';
import { Redirect } from 'wouter';
import UserScreenContainer from '../components/containers/UserScreenContainer';

const UserScreen = ({ userId }: { userId: string }) => {
	const { isLoading, isError } = useUserLogged(userId);

	if (isLoading) {
		return <p>Loading..</p>;
	}

	if (isError) {
		return <Redirect to='/' />;
	}

	return <UserScreenContainer />;
};

export default UserScreen;
