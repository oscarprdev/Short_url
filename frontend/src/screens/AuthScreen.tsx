import { useUserLogged } from '../hooks/useUserLogged';
import { ReactNode } from 'react';
import { Redirect } from 'wouter';

const AuthScreen = ({ children, userId }: { children: ReactNode; userId: string }) => {
	const { isLoading, isError } = useUserLogged(userId);

	if (isLoading) {
		return <p>Loading..</p>;
	}

	if (isError) {
		return <Redirect to="/" />;
	}

	return <>{children}</>;
};

export default AuthScreen;
