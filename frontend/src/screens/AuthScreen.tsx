import { ReactNode } from 'react';
import { useQueryParams } from '../hooks/useQueryParams';
import { useUserLogged } from '../hooks/useUserLogged';
import { Redirect } from 'wouter';

const AuthScreen = ({ children }: { children: ReactNode }) => {
	const { queryParams } = useQueryParams();

	const { isLoading, isError } = useUserLogged(queryParams.user);

	if (isLoading) {
		return <p>Loading..</p>;
	}

	if (isError) {
		return <Redirect to='/' />;
	}

	return <>{children}</>;
};

export default AuthScreen;
