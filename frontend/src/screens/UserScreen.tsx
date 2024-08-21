import UserScreenContainer from '../components/containers/UserScreenContainer';
import { useUserLogged } from '../hooks/useUserLogged';
import { IconLoader2 } from '@tabler/icons-react';
import { Redirect } from 'wouter';

const UserScreen = ({ userId }: { userId: string }) => {
	const { isLoading, isError } = useUserLogged(userId);

	if (isLoading) {
		return (
			<section className="grid place-items-center w-screen h-screen">
				<IconLoader2 className="animate-spin text-contrast" size={30} />
			</section>
		);
	}

	if (isError) {
		return <Redirect to="/" />;
	}

	return <UserScreenContainer />;
};

export default UserScreen;
