from typing import List
from app.models import MyTeam, Player


class UpdateMyTeamPlayers:
    def execute(self, my_team_uuid, players_uuid: List[int]):
        my_team = MyTeam.objects.get(uuid=my_team_uuid)
        players = Player.objects.filter(uuid__in=players_uuid)
        my_team.players.set(players)
