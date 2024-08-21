import { API_URL } from '../../constants/apiUrl';
import { useGlobalStore } from '../../store/globalState';
import { Button } from '../ui/button';
import { IconUser } from '@tabler/icons-react';

const ButtonLogout = () => {
	const { clearStore } = useGlobalStore();

	const onLogoutClick = () => {
		clearStore();
		window.location.href = `${API_URL}/auth/logout`;
	};
	return (
		<Button data-testid="logout-btn" onClick={onLogoutClick}>
			<IconUser />
			Logout
		</Button>
	);
};

export default ButtonLogout;
