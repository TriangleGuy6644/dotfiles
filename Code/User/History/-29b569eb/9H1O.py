import os
import discord
from discord import app_commands
from google.cloud import aiplatform
from dotenv import load_dotenv

load_dotenv()

DISCORD_TOKEN = os.getenv("DISCORD_TOKEN")
GUILD_ID = int(os.getenv("GUILD_ID"))
GEMINI_API_KEY = os.getenv("GEMINI_API_KEY")

intents = discord.Intents.default()
intents.message_content = True
client = discord.Client(intents=intents)
tree = app_commands.CommandTree(client)

aiplatform.init(api_key=GEMINI_API_KEY, project="YOUR_PROJECT_ID", location="us-central1")

@tree.command(name="summarize", description="Summarize the last [number] messages", guild=discord.Object(id=GUILD_ID))
@app_commands.describe(number="Number of messages to summarize (max 100)")
async def summarize(interaction: discord.Interaction, number: int = 100):
    if number > 100:
        number = 100
    channel = interaction.channel
    messages = await channel.history(limit=number).flatten()
    messages_text = "\n".join(msg.content for msg in reversed(messages) if msg.author != client.user)

    if not messages_text:
        await interaction.response.send_message("No messages to summarize")
        return

    response = aiplatform.TextGenerationModel.from_pretrained("gemini-2.5").predict(
        f"Summarize the following Discord messages:\n{messages_text}"
    )

    summary = response.text[:2000]  
    await interaction.response.send_message(f"**Summary:**\n{summary}")

@client.event
async def on_ready():
    await tree.sync()
    print(f"Logged in as {client.user}!")

client.run(DISCORD_TOKEN)
